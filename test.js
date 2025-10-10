import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate, Trend } from 'k6/metrics';

// Custom metrics
const errorRate = new Rate('error_rate');
const responseTime = new Trend('response_time');

// Test configuration
export const options = {
    stages: [
        { duration: '10s', target: 50 },   // Ramp up to 50 VUs in 10 seconds
        { duration: '50s', target: 100 },  // Maintain 100 VUs for 50 seconds
        { duration: '10s', target: 0 },    // Ramp down to 0
    ],
    thresholds: {
        'http_req_duration': ['p(95)<100'], // 95% of requests < 100ms
        'error_rate': ['rate<0.01'],        // Error rate < 1%
        'http_req_failed': ['rate<0.01'], // Failed requests < 1%
    },
};

// Setup
export function setup() {
    console.log('Starting K6 Performance Test');
    console.log('Target URL: http://localhost:3000/finenance/category');
}

// Main test
export default function () {
    const url = 'http://localhost:3000/finenance/category';
    const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNyIsImlzcyI6ImZpbmVuYW5jZS1iYWNrZW5kIiwiZXhwIjoxNzYwMDg1Nzc0LCJuYmYiOjE3NjAwODIxNzQsImlhdCI6MTc2MDA4MjE3NH0.phqCrjKGZegceeVO_prE0FvRy_rYIso2yN-n8bnk4M4';

    const params = {
        headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json',
        },
    };

    // Correct usage of params
    const response = http.get(url, params);

    // Record custom metrics
    responseTime.add(response.timings.duration);
    errorRate.add(response.status !== 200);

    // Checks
    const isSuccess = check(response, {
        'status is 200': (r) => r.status === 200,
        'response time < 100ms': (r) => r.timings.duration < 100,
        'response has body': (r) => r.body && r.body.length > 0,
    });

    if (!isSuccess) {
        console.error(`Request failed - Status: ${response.status}, Duration: ${response.timings.duration}ms`);
    }

    sleep(0.1); // optional small delay
}

// Teardown
export function teardown() {
    console.log('K6 Performance Test completed');
}

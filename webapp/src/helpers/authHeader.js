export function authHeader() {
    // return authorization header with basic auth credentials
    let token = localStorage.getItem('token');


    return {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Headers": "Origin, Content-Type, X-Auth-Token",
        "Access-Control-Allow-Methods": "GET",
        "Content-Type": "application/json;charset=UTF-8"
    }
}

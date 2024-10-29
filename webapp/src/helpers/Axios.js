import axios from 'axios';
import { authHeader } from './authHeader';
import https from 'https';

const Axios = axios.create({
    baseURL: "https://localhost/api",
    timeout: 3000,
    headers: authHeader(),
    httpsAgent: new https.Agent({  
        rejectUnauthorized: false
      })
});

export default Axios;

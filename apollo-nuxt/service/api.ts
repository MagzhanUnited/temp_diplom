import axios from 'axios';
import { RESIDENT_API } from '../config';
const instance = axios.create({ baseURL: String(RESIDENT_API), withXSRFToken: true, withCredentials: true });
export default instance;

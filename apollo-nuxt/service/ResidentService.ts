import { getHeader } from '../config';
import api from './api';
const ResidentService = {
    async getResidents(category: string) {
        return await api.get(`/residentialas/${category}/undefined`, { headers: getHeader() });
    },
    async postResident(request: any) {
        console.log('request:', request);

        return await api.post('/residential', request, { headers: getHeader() });
    },
    async deleteResident(id: any) {
        console.log('A:', id);

        await api.delete(`/residential/${id}`, {
            headers: getHeader()
        });
    },
    async upload(formData: any) {
        await api.post('/upload', formData, {
            headers: {
                'Access-Control-Allow-Credentials': 'true',
                'Content-Type': 'multipart/form-data',
                Accept: '*/*'
            }
        });
    },
    async login(request: any) {
        return await api.post('/auth', request, { headers: getHeader() });
    }
};

export default ResidentService;

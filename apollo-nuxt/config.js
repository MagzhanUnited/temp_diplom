export const RESIDENT_API = import.meta.env.VITE_APP_API_URL;
//"http://10.255.184.3:5001";
export const getHeader = function (access_token = '') {
    return {
        'Access-Control-Allow-Credentials': 'true',
        'Content-Type': 'application/json',
        Accept: '*/*'
    };
};

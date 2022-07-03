import axios from 'axios';


export const API_URL_ENV = {
    DEV: {
        BASE_URL: '',
        LIFF_ID: '1657215266-jJxNVnRl',
        LINE_CLIENT_ID: '1657215266'

    }

}
export const API_URL = API_URL_ENV.DEV;

export const ROLES = {
    'Public': 10000,
    'User': 10001,
    'Vip': 10002,
    'Admin': 10003
}
export const GET_ROLES = (n) => {
    if (ROLES.Public === n) {
        return 'Public'
    }
    if (ROLES.User === n) {
        return 'User'
    }
    if (ROLES.Vip === n) {
        return 'Vip'
    }
    if (ROLES.Admin === n) {
        return 'Admin'
    }
    return ''
}

export let accessToken;
export const setToken = (token) => {
    accessToken = token;
};

export const getToken = () => {
    return accessToken;
};

export const HEADER = {
    "headers": {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE',
        'Access-Control-Allow-Headers': '*',
        'Accept': '*/*'
    }
}

export const GET = (endpoint) => {
    return axios.get(`${API_URL.BASE_URL}${endpoint}`, HEADER);
}
export const POST = (endpoint, apiPayload) => {
    return axios.post(`${API_URL.BASE_URL}${endpoint}`, apiPayload, HEADER);
}

export let LineUserIdToken 
export let LineUserAccessToken 
export const setLineToken = (lineUserIdToken,lineUserAccessToken) => {
    LineUserIdToken = lineUserIdToken
    LineUserAccessToken = lineUserAccessToken
    console.log("setLineToken : ",LineUserAccessToken)
}


export default {
    ROLES,
    GET_ROLES,
    GET,
    POST,
    getToken,
    setToken

};
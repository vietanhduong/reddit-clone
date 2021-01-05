import axios from "axios";
const API = process.env.REACT_APP_API || "/api";
const md5 = require('md5');

export const fetch = async () => {
    const res = await axios.get(`${API}/topics`);
    return res.data;
};

export const login = async (username, password) => {
    const payload = {
        "username": username,
        "password": md5(password || "")
    }
    const res = await axios.post(`${API}/login`, payload);
    const data = res.data;
    return data;
}


export const me = async () => {
    const accessToken = localStorage.getItem("access_token");
    if (!accessToken) {
        return;
    } 
    const config = {
        headers: {
            Authorization: `Bearer ${accessToken}`
        }
    }
    const res = await axios.get(`${API}/users/me`, config);
    const data = res.data;
    return data;
}

export const vote = async (id, isUp) => {
    const accessToken = localStorage.getItem("access_token");
    if (!accessToken) {
        return alert("please login first");
    } 
    const config = {
        headers: {
            Authorization: `Bearer ${accessToken}`
        }
    }
    const res = await axios.post(`${API}/topics/${id}/${isUp?"upvote":"downvote"}`, {}, config);
    const data = res.data;
    return data;
} 


export const create = async (content) => {
    const accessToken = localStorage.getItem("access_token");
    if (!accessToken) {
        return alert("please login first");
    } 
    const config = {
        headers: {
            Authorization: `Bearer ${accessToken}`
        }
    }
    const res = await axios.post(`${API}/topics`, {"content": content}, config);
    const data = res.data;
    return data;
} 
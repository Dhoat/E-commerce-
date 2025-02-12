import axios from "axios";

export const register = (userData) => {
    return axios.post("http://localhost:8080/api/register", userData, {
        headers: {
            "Content-Type": "application/json"
        }
    });
};


export const login = (userData) => {
    return axios.post("http://localhost:8080/api/login", userData, {
        headers: {
            "Content-Type": "application/json"
        }
    });
};

// Add product

export const addProduct = (userData) => {
    return axios.post("http://localhost:8080/api/products", userData, {
        headers: {
            "Content-Type": "application/json"
        }
    });
};



export const fetchProducts = async () => {
    return await axios.get("http://localhost:8080/api/products", {
        headers: {
            "Content-Type": "application/json",
           
        }
    });
};


export const shortenURL = async (url) => {
    return axios.post("http://localhost:8080/api/shorten", { url }, {
        headers: {
            "Content-Type": "application/json"
        }
    });
};

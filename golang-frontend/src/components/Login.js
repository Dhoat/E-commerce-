// src/components/Login.js

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../apiService";
import { useAuth } from "../AuthContext";
import "./Register.css"; // Reuse the same CSS file for styling

const Login = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [message, setMessage] = useState("");
    const navigate = useNavigate();
    const { login: authLogin } = useAuth();

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const response = await login({ username, password });
            console.log("Login successful:", response.data);
            setMessage("Login successful");

            // Authenticate the user
            authLogin({ username: response.data.username });
            console.log("User authenticated:", { username: response.data.username });

            // Redirect to the new page
            navigate('/dashboard'); // Replace '/dashboard' with the desired route
        } catch (error) {
            console.error("Login failed:", error.response ? error.response.data : error.message);
            setMessage(error.response ? error.response.data.error : "Login failed");
        }
    };

    return (
        <div className="form-container">
            <form className="form" onSubmit={handleSubmit}>
                <input
                    type="text"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    placeholder="Username"
                    required
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Password"
                    required
                />
                <button type="submit">Login</button>
                {message && <p className="message">{message}</p>}
            </form>
        </div>
    );
};

export default Login;

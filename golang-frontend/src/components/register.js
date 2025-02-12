// src/components/Register.js

import React, { useState } from "react";
import { register } from "../apiService";
import "./Register.css";

const Register = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [email, setEmail] = useState("");
    const [message, setMessage] = useState("");

    const handleSubmit = async (event) => {
        event.preventDefault();
        try {
            const response = await register({ username, password, email });
            console.log("Registration successful:", response.data);
            setMessage("Registration successful");
        } catch (error) {
            console.error("Registration failed:", error.response ? error.response.data : error.message);
            setMessage(error.response ? error.response.data.error : "Registration failed");
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
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="Email"
                    required
                />
                <button type="submit">Register</button>
                {message && <p className="message">{message}</p>}
            </form>
        </div>
    );
};

export default Register;

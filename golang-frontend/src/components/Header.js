// src/components/Header.js

import React from "react";
import { useAuth } from "../AuthContext";
import "./Layout.css";

const Header = () => {
    const { isAuthenticated, user, logout } = useAuth();
    console.log("Header render - isAuthenticated:", isAuthenticated, "user:", user);
    return (
        <header className="header">
            <div className="header-logo">FlipKart Clone</div>
            <nav className="header-nav">
                {isAuthenticated ? (
                    <>
                        <a href="/profile">{user?.username || "Profile"}</a>
                        <a href="/" onClick={logout}>Logout</a>
                    </>
                ) : (
                    <>
                        <a href="/login">Login</a>
                        <a href="/register">Register</a>
                    </>
                )}
            </nav>
        </header>
    );
};

export default Header;

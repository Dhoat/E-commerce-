// src/AuthContext.js

import React, { createContext, useState, useContext, useEffect } from "react";

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(() => {
        return !!sessionStorage.getItem('isAuthenticated');
    });
    const [user, setUser] = useState(() => {
        return JSON.parse(sessionStorage.getItem('user'));
    });

    const login = (userData) => {
        setIsAuthenticated(true);
        setUser(userData);
        sessionStorage.setItem('isAuthenticated', 'true');
        sessionStorage.setItem('user', JSON.stringify(userData));
        console.log("User logged in:", userData);
    };

    const logout = () => {
        setIsAuthenticated(false);
        setUser(null);
        sessionStorage.removeItem('isAuthenticated');
        sessionStorage.removeItem('user');
        console.log("User logged out");
    };

    useEffect(() => {
        console.log("Auth state changed:", { isAuthenticated, user });
    }, [isAuthenticated, user]);

    return (
        <AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => useContext(AuthContext);

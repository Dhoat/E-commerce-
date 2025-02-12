// src/components/Dashboard.js

import React from "react";
import Header from './Header';
import Footer from './Footer'; // Create a CSS file for styling this component
import { AuthProvider } from '../AuthContext';
console.log(AuthProvider)
const Dashboard = () => {
    return (
        <AuthProvider>

        <div className="dashboard-container">
            <Header />
            <h1>Welcome to your Dashboard!</h1>
            <p>This is a protected route that you can access after logging in.</p>
            <Footer />
        </div>
        </AuthProvider>
    );
};

export default Dashboard;

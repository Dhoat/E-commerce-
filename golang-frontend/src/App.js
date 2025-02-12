import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import './App.css';
import Register from '../src/components/register';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import { AuthProvider } from './AuthContext';
import Product from './components/Product';
import ProductList from './components/ProductList';
import URLShortener from './components/URLShortener';

function App() {
    return (
      <AuthProvider>
        <Router>
            <Routes>
                <Route path="/register" element={<Register />} />
                <Route path="/" element={<Login />} />
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/products" element={<Product />} />
                <Route path="/product-list" element={<ProductList />} />
                <Route path="/shorten-url" element={<URLShortener />} />
            </Routes>
        </Router>
      </AuthProvider>
    );
}

export default App;

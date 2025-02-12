import React, { useState } from 'react';
import { addProduct } from '../apiService';
import './Product.css';
import Header from './Header';  // Make sure to import the Header component
import Footer from './Footer';  // Make sure to import the Footer component

const Product = () => {
    const [product, setProduct] = useState({
        name: '',
        description: '',
        price: 0,
        stock: 0
    });
    const [message, setMessage] = useState('');

    const handleChange = (e) => {
        const { name, value } = e.target;
        setProduct({
            ...product,
            [name]: name === 'price' || name === 'stock' ? parseFloat(value) : value
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log("Submitting product:", product); // Log the product payload
        try {
            const response = await addProduct(product);
            console.log("Product added successfully:", response.data);
            setMessage("Product added successfully");
            setProduct({
                name: '',
                description: '',
                price: 0,
                stock: 0
            });
        } catch (error) {
            console.error("Failed to add product:", error.response ? error.response.data : error.message);
            setMessage(error.response ? error.response.data.error : "Failed to add product");
        }
    };

    return (
        <>
            <Header /> {/* Add Header component */}
            <div className="form-container">
                <form className="form" onSubmit={handleSubmit}>
                    <input
                        type="text"
                        name="name"
                        value={product.name}
                        onChange={handleChange}
                        placeholder="Name"
                        required
                    />
                    <input
                        type="text"
                        name="description"
                        value={product.description}
                        onChange={handleChange}
                        placeholder="Description"
                        required
                    />
                    <input
                        type="number"
                        name="price"
                        value={product.price}
                        onChange={handleChange}
                        placeholder="Price"
                        required
                    />
                    <input
                        type="number"
                        name="stock"
                        value={product.stock}
                        onChange={handleChange}
                        placeholder="Stock"
                        required
                    />
                    <button type="submit">Add Product</button>
                    {message && <p className="message">{message}</p>}
                </form>
            </div>
            <Footer /> {/* Add Footer component */}
        </>
    );
};

export default Product;

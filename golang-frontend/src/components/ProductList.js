// src/components/ProductList.js
import React, { useEffect, useState } from 'react';
import { fetchProducts } from '../apiService';
import ProductCard from './ProductCard';


const ProductList = () => {
    const [products, setProducts] = useState([]);
    const [error, setError] = useState(null);

    useEffect(() => {
        const getProducts = async () => {
            try {
                const response = await fetchProducts();
                setProducts(response.data);
            } catch (error) {
                setError("Failed to fetch products");
            }
        };

        getProducts();
    }, []);

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <div className="product-list">
            {products.map((product) => (
                <ProductCard key={product.id} product={product} />
            ))}
        </div>
    );
};

export default ProductList;

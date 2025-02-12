import React, { useState } from "react";
import { shortenURL } from "../apiService";

const URLShortener = () => {
    const [originalURL, setOriginalURL] = useState("");
    const [shortURL, setShortURL] = useState("");
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const handleShorten = async () => {
        if (!originalURL) {
            setError("Please enter a valid URL");
            return;
        }

        setLoading(true);
        setError("");
        try {
            const response = await shortenURL(originalURL);
            setShortURL(`http://localhost:8080/${response.data.short_url}`);
        } catch (error) {
            setError(error.response ? error.response.data.error : "Failed to shorten URL");
        }
        setLoading(false);
    };

    return (
        <div className="form-container">
            <h2>URL Shortener</h2>
            <input
                type="text"
                value={originalURL}
                onChange={(e) => setOriginalURL(e.target.value)}
                placeholder="Enter URL..."
            />
            <button onClick={handleShorten} disabled={loading}>
                {loading ? "Shortening..." : "Shorten URL"}
            </button>
            {error && <p className="error">{error}</p>}
            {shortURL && (
                <p>
                    Shortened URL: <a href={shortURL} target="_blank" rel="noopener noreferrer">{shortURL}</a>
                </p>
            )}
        </div>
    );
};

export default URLShortener;

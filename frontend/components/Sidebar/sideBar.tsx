import React from 'react';
import styles from './sideBar.module.css';
import { useState } from 'react';

export default function Sidebar() {

    const [minPrice, setMinPrice] = useState('');
    const [maxPrice, setMaxPrice] = useState('');

    const handleMinPriceChange = (event) => {
        setMinPrice(event.target.value);
    };

    const handleMaxPriceChange = (event) => {
        setMaxPrice(event.target.value);
    };

    return (
        <div className="sidebar">
            <h2>Filter</h2>

            <div className="filter">
                <label htmlFor="make">Make:</label>
                <select id="make" name="make">
                    <option value="all">All Makes</option>
                    <option value="toyota">Toyota</option>
                    <option value="tesla">Tesla</option>
                    {/* Add more options as needed */}
                </select>
            </div>

            <div className="filter">
                <label htmlFor="model">Model:</label>
                <select id="model" name="model">
                    <option value="all">All Models</option>
                    <option value="corolla">Corolla</option>
                    <option value="model-s">Model S</option>
                    {/* Add more options as needed */}
                </select>
            </div>

            <div className="filter">
                <label htmlFor="trim">Trim:</label>
                <select id="trim" name="trim">
                    <option value="all">All trims</option>
                    <option value="SE">SE</option>
                    <option value="L">L</option>
                    {/* Add more options as needed */}
                </select>
            </div>

            <h4 className={styles.heading}>Price</h4>

            <div className={styles.filterOption}>
                <label htmlFor="minPrice">Min Price:</label>
                <input
                    type="number"
                    id="minPrice"
                    name="minPrice"
                    value={minPrice}
                    onChange={handleMinPriceChange}
                />
            </div>

            <div className={styles.filterOption}>
                <label htmlFor="maxPrice">Max Price:</label>
                <input
                    type="number"
                    id="maxPrice"
                    name="maxPrice"
                    value={maxPrice}
                    onChange={handleMaxPriceChange}
                />
            </div>

        </div>
    );
}

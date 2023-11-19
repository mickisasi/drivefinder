import React from 'react';
import './sideBar.module.css';

export default function Sidebar() {
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

    </div>
  );
}

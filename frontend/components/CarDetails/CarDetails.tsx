// CarDetails.tsx
import React, { useState, useEffect } from 'react';
import styles from './styles.module.css';

interface CarDetailsProps {
    make: string;
}

const CarDetails: React.FC<CarDetailsProps> = ({ make }) => {
    const [title, setTitle] = useState('');
    const [description, setDescription] = useState('');

    const updateCarDetails = (selectedMake: string) => {
        // Logic to fetch details based on the selected make
        // For simplicity, I'm using static content here
        if (selectedMake === 'Toyota' || selectedMake === 'toyota') {
            setTitle('Toyota Corolla');
            setDescription(
                'The Toyota Corolla is one of the best-selling cars globally, and...'
            );
        } else if (selectedMake === 'Dodge' || selectedMake === 'dodge') {
            setTitle('Dodge Charger');
            setDescription(
                'The Dodge Charger is a high-performance sedan known for its powerful engine and...'
            );
        }
        // Add more cases for other makes as needed
    };

    useEffect(() => {
        updateCarDetails(make);
    }, [make]);

    return (
        <div className={styles.descriptionContainer}>
            <h3>{title}</h3>
            <p>{description}</p>
        </div>
    );
};

export default CarDetails;

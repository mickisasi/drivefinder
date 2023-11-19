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
            setTitle('Toyota');
            setDescription(
                'The benefits of purchasing a vehicle from Toyota are plentiful. The manufacturer offers a wide array of automobile options covering nearly every consumers needs and budget. The Toyota collection includes rugged off-roaders, small urban-focused hatchbacks, and everything in between. An often-overlooked bonus of purchasing a vehicle from one of the world’s largest producers is dependability.'
            );
        } else if (selectedMake === 'Dodge' || selectedMake === 'dodge') {
            setTitle('Dodge');
            setDescription(
                'The Dodge Charger is a high-performance sedan known for its powerful engine and...'
            );
        } else if (selectedMake === 'Tesla' || selectedMake === 'tesla') {
            setTitle('Tesla');
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

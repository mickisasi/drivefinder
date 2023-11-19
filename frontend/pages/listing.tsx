// listing.tsx

import { FeaturesCard } from '@/components/FeaturesCard/FeaturesCard';
import { HeaderSimple } from '@/components/Header/Header';
import { MantineProvider } from '@mantine/core';

import React from 'react';
import styles from './styles.module.css';

export default function Listing() {
  return (
    <MantineProvider>
      <HeaderSimple />
      <div className={styles.container}>
        <div className={styles.descriptionContainer}>
          <h3>Tesla Model S</h3>
          <p>
            The Toyota Corolla is one of the best-selling cars globally, and at times has been the definitive leader in sales.
            Surpassing the Volkswagen Beetle in 1997, the Toyota Corolla is the best-selling nametag in history.
            As of 2021, there have been over 50 million Toyota Corollas manufactured globally.
            Inarguably, the sedan has set the standard for subcompact and compact cars since its introduction in 1966.
            Currently, in its twelfth generation, the Toyota Corolla continues to be a tried, tested, and proper vehicle
            with production locations in all corners of the globe.
          </p>
        </div>
        <h2>Car Listings</h2>
        <div className={styles.cardList}>
          {[...Array(5)].map((_, index) => (
            <FeaturesCard key={index} />
          ))}
        </div>
      </div>
    </MantineProvider>
  );
}
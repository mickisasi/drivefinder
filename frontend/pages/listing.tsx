import { FeaturesCard } from '@/components/FeaturesCard/FeaturesCard'
import { MantineProvider } from '@mantine/core';

import React from 'react'

export default function listing() {
  return (
    <MantineProvider>
        <div>
            <h2>Car Listings</h2>
            <FeaturesCard />
        </div>
    </MantineProvider>
  )
}

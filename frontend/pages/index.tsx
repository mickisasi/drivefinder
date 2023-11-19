import React from 'react'
import Image from 'next/image'
import { Inter } from 'next/font/google'
const inter = Inter({ subsets: ['latin'] })
import { HeaderSimple } from '@/components/Header/Header'
import { FilterCarsDropdown } from '@/components/FilterCarsDropdown/FilterCarsDropdown'
import { HeroImageBackground } from '@/components/Hero/Hero'
import { Footer } from '@/components/Footer/footer'
export default function Home() {
  return (
    <main

    >
    {/* <HeaderSimple/>
    <FilterCarsDropdown/> */}
    <HeroImageBackground/>
    <Footer/>
    </main>

  )
}
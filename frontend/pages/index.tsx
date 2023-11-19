import React from 'react'
import Image from 'next/image'
import { Inter } from 'next/font/google'
const inter = Inter({ subsets: ['latin'] })
import { HeaderSimple } from '@/components/Header/Header'

  export default function Home() {
    
    return (
      <main
      >
      <HeaderSimple/>
      </main>
    );
}
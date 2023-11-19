import React from 'react'
import styles from './styles.module.css'
import { FaUserCircle } from "react-icons/fa";
// import font


export default function Navbar() {
  return (
    <div className={styles.innerContainer}>
    <header>
        <nav>
            <h1>Drive Finder</h1>
        </nav>
    </header>
    </div>
  )
}

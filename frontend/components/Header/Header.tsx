import React from 'react'
import styles from './styles.module.css'
import { Container, Group, Burger } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import Title from '../Title/Title';
import classes from './styles.module.css';
import { useState } from 'react';
import Button from 'react'


const links = [
  { link: '/Home', label: 'Home' },
];

export function HeaderSimple() {
  const [opened, { toggle }] = useDisclosure(false);
  const [active, setActive] = useState(links[0].link);

  const items = links.map((link) => (
    <a
      key={link.label}
      href={link.link}
      className={classes.link}
      data-active={active === link.link || undefined}
      onClick={(event) => {
        event.preventDefault();
        setActive(link.link);
      }}
    >
      {link.label}
    </a>
  ));

  return (
    <div className={styles.innerContainer}>
    <header className={classes.header}>
      <Container size="md" className={classes.inner}>
        <Title/>
        <Group gap={5} visibleFrom="xs">
          <div className={styles.order}>
               <svg fill="none" width="24px" height="24px" stroke="currentColor" stroke-width="1" class="w-4 h-4" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" d="M2.25 12l8.954-8.955a1.126 1.126 0 011.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25"/></svg>
          {items}
          </div>
        </Group>
        {/* <Group visibleFrom="sm">
            <Button variant="default">Log in</Button>
            <Button>Sign up</Button>
          </Group> */}
        <Burger opened={opened} onClick={toggle} hiddenFrom="xs" size="sm" />
      </Container>
    </header>
    </div> 
  );
}
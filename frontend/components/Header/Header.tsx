import React from 'react'
import styles from './styles.module.css'
import { Container, Group, Burger } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import Title from '../Title/Title';
import classes from './styles.module.css';
import { useState } from 'react';

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
          {items}
        </Group>
        <Burger opened={opened} onClick={toggle} hiddenFrom="xs" size="sm" />
      </Container>
    </header>
    </div> 
  );
}
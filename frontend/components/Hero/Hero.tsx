import cx from 'clsx';
import { Title, Text, Container, Button, Overlay } from '@mantine/core';
import classes from './styles.module.css';
import { FilterCarsDropdown } from '../FilterCarsDropdown/FilterCarsDropdown';

export function HeroImageBackground() {
  return (
    <div className={classes.wrapper}>
      <Overlay color="#000" opacity={0.65} zIndex={1} />
      <div className={classes.inner}>
        <Title className={classes.title}>
            Drive Finder
        </Title>
        <Container size={640}>
          <Text size="lg" className={classes.description}>
          Embark on this exciting journey to create a revolutionary solution that enhances the car buying experience through intelligent automation and personalized notifications.
          </Text>
        </Container>
        <div className={classes.controls}>
         <FilterCarsDropdown/>
        </div>
      </div>
    </div>
  );
}
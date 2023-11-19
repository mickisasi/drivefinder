import cx from "clsx";
import {
  Title,
  Text,
  Container,
  Button,
  Overlay,
  Flex,
  Avatar,
  Modal,
  TextInput,
  PasswordInput,
} from "@mantine/core";
import classes from "./styles.module.css";
import { FilterCarsDropdown } from "../FilterCarsDropdown/FilterCarsDropdown";
import { useState } from "react";

export function HeroImageBackground() {
  const [opened, setOpened] = useState(false);

  return (
    <>
      <Modal
        opened={opened}
        onClose={() => setOpened(false)}
        size={"lg"}
        overlayProps={{ backgroundOpacity: 0.5, blur: 0.5 }}
      >
        <div style={{ padding: "1rem" }}>
          <TextInput label="Email" placeholder="Email" />
          <PasswordInput label="Password" placeholder="Password" />
        </div>
      </Modal>
      <button onClick={() => setOpened(true)}>click me</button>
      <div className={classes.wrapper}>
        <Overlay color="#000" opacity={0.65} zIndex={1} />
        <div className={classes.inner}>
          <Flex>
            <Title className={classes.title}>Drive Finder</Title>
            // <Avatar variant="transparent" radius="xl" size="xl" src="" />;
          </Flex>
          <Container size={640}>
            <Text size="lg" className={classes.description}>
              Embark on this exciting journey to create a revolutionary solution
              that enhances the car buying experience through intelligent
              automation and personalized notifications.
            </Text>
          </Container>
          <div className={classes.controls}>
            <FilterCarsDropdown />
          </div>
        </div>
      </div>
    </>
  );
}

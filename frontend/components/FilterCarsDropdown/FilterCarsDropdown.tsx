import React from "react";
import { Select, TextInput, Grid, Container } from "@mantine/core";
import classes from "./styles.module.css";

export function FilterCarsDropdown() {
  return (
    <Container size="md">
      <div className={classes.card}>
        <Grid>
          <Grid.Col span={12} md={6}>
            <div className={classes.card}>
              <TextInput
                placeholder="Any Make"
                classNames={classes}
              />
              <Select
                mt="md"
                comboboxProps={{ withinPortal: true }}
                data={["React", "Angular", "Svelte", "Vue"]}
                placeholder="Pick one"
                classNames={classes}
              />
            </div>
          </Grid.Col>

          <Grid.Col span={12} md={6}>
            <div className={classes.card}>
              <Grid>
                <Grid.Col span={6}>
                  <TextInput
                    placeholder="Enter Postal Code"
                    classNames={classes}
                  />
                </Grid.Col>
                <Grid.Col span={6}>
                  <TextInput
                    placeholder="Enter Price"
                    classNames={classes}
                  />
                </Grid.Col>
              </Grid>
            </div>
          </Grid.Col>
        </Grid>
      </div>
    </Container>
  );
}

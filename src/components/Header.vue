<template>
  <header>
    <Button text="Tasks" color="black" />
    <Button text="Profiles" color="black" />
    <Button text="Cards" color="black" />
  </header>
</template>

<script>
import Button from "./Button";

export default {
  name: "Header-top",
  props: {
    title: String,
  },
  components: {
    Button,
  },
};
</script>

<style scoped>
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>

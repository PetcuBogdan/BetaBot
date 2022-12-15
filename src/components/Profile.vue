<template>
  <option value="profile.fname">
    {{ profile.fname }}
  </option>
</template>

<script>
export default {
  name: "Profile-el",
  props: {
    profile: Object,
  },
};
</script>

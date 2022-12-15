<template>
  <div>
    <select>
      <option value="0" label="modifyProfile" selected class="selectDefault">
        Modify profile ...
      </option>
    </select>
    <select :key="profile.id" v-for="profile in profiles">
      <Profile :profile="profile" />
    </select>
  </div>
</template>

<script>
import Profile from "./Profile";

export default {
  name: "Profiles-bot",
  props: {
    profiles: Array,
  },
  components: {
    Profile,
  },
};
</script>

<template>
  <button :style="{ background: color }" class="btn">
    {{ text }}
  </button>
</template>

<script>
export default {
  name: "Button-top",
  props: {
    text: String,
    color: String,
  },
};
</script>

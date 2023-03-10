<template>
  <div id="app">
    <nav-bar />
    <router-view />
    <footer-bar />
  </div>
</template>

<script>
import NavBar from "@/components/NavBar.vue";
import FooterBar from "@/components/FooterBar.vue";

export default {
  name: "app",
  components: {
    NavBar,
    FooterBar,
  },
  data() {
    return {};
  },
  created() {
    this.$store.commit("user", {
      name: "John Doe",
      email: "john@example.com",
     
    });
  },
};
</script>

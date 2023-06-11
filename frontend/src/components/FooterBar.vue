<template>
  <footer v-show="isFooterBarVisible" class="footer">
    <div class="container-fluid">
      <div class="level">
        <div class="level-left">
          <div class="level-item">
            <div class="footer-copyright">
              <b>&copy; {{ year }}, Sistema Integral de Salud Materna</b>.
            </div>
          </div>
        </div>
        <div class="level-center">
          <div class="level-item">
            <div class="logo">
              <img src="../assets/secretaria.png" />
              <img src="../assets/tec.png" />
              <img src="../assets/materno.png" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "FooterBar",
  data() {
    return {
      year: new Date().getFullYear(),
    };
  },
  computed: {
    ...mapState(["isFooterBarVisible"]),
  },
};
</script>
<style>
.logo img {
height: fit-content;
width:10%;
aspect-ratio: 3/2;
object-fit: contain;

}
</style>

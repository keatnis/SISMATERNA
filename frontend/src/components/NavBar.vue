<template>
  <nav v-show="isNavBarVisible" id="navbar-main" class="navbar is-fixed-top">
    <div class="navbar-brand">
      <a
        class="navbar-item is-hidden-desktop"
        @click.prevent="asideToggleMobile"
      >
        <b-icon :icon="asideMobileIcon" />
      </a>
      <a
        class="navbar-item is-hidden-touch is-hidden-widescreen is-desktop-icon-only"
        @click.prevent="asideDesktopOnlyToggle"
      >
        <b-icon icon="menu" />
      </a>
    </div>
    <div class="navbar-brand is-right">
      <a
        class="navbar-item navbar-item-menu-toggle is-hidden-desktop"
        @click.prevent="menuToggle"
      >
        <b-icon :icon="menuToggleIcon" custom-size="default" />
      </a>
    </div>
    <div
      class="navbar-menu fadeIn animated faster"
      :class="{ 'is-active': isMenuActive }"
    >
      <div class="navbar-end">
        <nav-bar-menu class="has-divider">
          <b-icon icon="menu" custom-size="default" />
          <span>Menu</span>
          <div slot="dropdown" class="navbar-dropdown">
            <router-link
              to="/agregar-embarazada"
              class="navbar-item"
              exact-active-class="is-active"
            >
              <b-icon icon="account" custom-size="default" />
              <span>Embarazada</span>
            </router-link>
            <router-link
              to="/ListaEmbarazada"
              class="navbar-item"
              exact-active-class="is-active"
            >
              <b-icon icon="account" custom-size="default" />
              <span>Lista</span>
            </router-link>
            <a class="navbar-item">
              <b-icon icon="settings" custom-size="default" />
              <span>Lista</span>
            </a>
            <a class="navbar-item">
              <b-icon icon="email" custom-size="default" />
              <span>Messages</span>
            </a>
            <hr class="navbar-divider" />
            <a class="navbar-item">
              <router-link
              to="/ListUSers"
              class="navbar-item"
              exact-active-class="is-active"
            >
              <b-icon icon="logout" custom-size="default" />
              </router-link>
            </a>
          </div>
        </nav-bar-menu>
        <nav-bar-menu class="has-divider has-user-avatar">
          <div class="is-user-name">
            <span>{{ userName }}</span>
          </div>

          <div slot="dropdown" class="navbar-dropdown">
            <router-link
              to="/profile"
              class="navbar-item"
              exact-active-class="is-active"
            >
              <b-icon icon="account" custom-size="default" />
              <span>My Perfil</span>
            </router-link>
            <a class="navbar-item">
              <b-icon icon="settings" custom-size="default" />
              <span>Configuraci[on]</span>
            </a>
            <a class="navbar-item">
              <b-icon icon="email" custom-size="default" />
              <span>Notificaciones</span>
            </a>
            <hr class="navbar-divider" />
            <a class="navbar-item">
              <b-icon icon="logout" custom-size="default" />
              <span>Salir</span>
            </a>
          </div>
        </nav-bar-menu>

        <a
          class="navbar-item is-desktop-icon-only"
          title="Log out"
          @click="logout"
        >
          <b-icon icon="logout" custom-size="default" />
          <span>Log out</span>
        </a>
      </div>
    </div>
  </nav>
</template>

<script>
import { mapState } from "vuex";
import NavBarMenu from "@/components/NavBarMenu.vue";

export default {
  name: "NavBar",
  components: {
    NavBarMenu,
  },
  data() {
    return {
      isMenuActive: false,
    };
  },
  computed: {
    asideMobileIcon() {
      return this.isAsideMobileExpanded ? "backburger" : "forwardburger";
    },
    menuToggleIcon() {
      return this.isMenuActive ? "close" : "dots-vertical";
    },
    ...mapState(["isAsideMobileExpanded", "isNavBarVisible", "userName"]),
  },
  mounted() {
    this.$router.afterEach(() => {
      this.isMenuActive = false;
    });
  },
  methods: {
    asideToggleMobile() {
      this.$store.commit("asideMobileStateToggle");
    },
    asideDesktopOnlyToggle() {
      this.$store.dispatch("asideDesktopOnlyToggle");
    },
    menuToggle() {
      this.isMenuActive = !this.isMenuActive;
    },
    logout() {
      this.$buefy.snackbar.open({
        message: "Log out clicked",
        queue: false,
      });
    },
  },
};
</script>

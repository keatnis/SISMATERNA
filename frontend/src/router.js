import Vue from "vue";
import Router from "vue-router";

import RegistrarEmbarazada from "@/components/RegistrarEmbarazada";

import CreateUsers from "@/components/CreateUsers";

import RegistrarPuerpera from "@/components/RegistrarPuerpera";
import ActualizarEmbarazada from "@/components/ActualizarEmbarazada";
import ActualizarPuerperia from "@/components/ActualizarPuerperia";
import ListaPuerpera from "@/components/ListaPuerpera";
import ListaEmbarazada from "@/components/ListaEmbarazada";

import Tarjetas from "@/components/Tarjetas";


import Login from '@/components/Mujer'

import ListaUsuario from '@/components/ListaUsuario'
Vue.use(Router);
const routes = [
  {
    path: "/full-page",
    component: () =>
      import(/* webpackChunkName: "full-page" */ "@/views/FullPage.vue"),
    children: [
      {
        meta: {
          title: "Login",
        },
        path: "/login",
        name: "login",
        component: () =>
          import(
            /* webpackChunkName: "full-page" */ "@/views/full-page/Login.vue"
          ),
      },
    ],
  },
  {
    meta: {
      title: 'Forms'
    },
    path: '/',
    name: 'Mujer',
    component: () =>
      import(/* webpackChunkName: "forms" */ '@/views/Tables.vue')
  },
  {
    meta: {
      title: 'Embarazada'
    },
    path: '/AddPregnant',
    name: 'Embarazada',
    component: () =>
      import(/* webpackChunkName: "forms" */ '@/views/AddPregnant.vue')
  },
  {
    path: "/agregar-embarazada",
    name: "RegistrarEmbarazada",
    component: RegistrarEmbarazada,
  },
  {
    path: "/Login2",
    name: "Lgin",
    component: Login,
  },
  {
    path: "/AddUser",
    name: "AddUser",
    component: CreateUsers,
  },

  {
    path: "/RegistrarPuerpera",
    name: "RegistrarPuerpera",
    component: RegistrarPuerpera,
  },
  {
    path: "/ActualizarEmbarazada",
    name: "ActualizarEmbarazada",
    component: ActualizarEmbarazada,
  },
  {
    path: "/ActualizarPuerperia",
    name: "ActualizarPuerperia",
    component: ActualizarPuerperia,
  },
  {
    path: "/ListaEmbarazada",
    name: "ListaEmbarazada",
    component: ListaEmbarazada,
  },
  {
    path: "/ListaPuerpera",
    name: "ListaPuerpera",
    component: ListaPuerpera,
  },

  {
    path: "/Tarjetas",
    name: "Tarjetas",
    component: Tarjetas,
  },
  {
    path: "/ListUsers",
    name: "USers",
    component: ListaUsuario,
  },
];
const router = new Router({
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
      
    } else {
      return { x: 0, y: 0 };
    }
    
  },
});

export default router;

export function useRouter() {
  return router;
}

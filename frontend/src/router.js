import Vue from 'vue'
import Router from 'vue-router'
import Mujer from '@/components/Mujer'
import RegistrarEmbarazada from '@/components/RegistrarEmbarazada'


import Login from '@/components/Login'
import RegistrarPuerpera from '@/components/RegistrarPuerpera'
import ActualizarEmbarazada from '@/components/ActualizarEmbarazada'
import ActualizarPuerperia from '@/components/ActualizarPuerperia'
import ListaPuerpera from '@/components/ListaPuerpera'
import ListaEmbarazada from '@/components/ListaEmbarazada'
import Informes from '@/components/Informes'
import Tarjetas from '@/components/Tarjetas'
import ListaUsuario from '@/components/ListaUsuario'
Vue.use(Router);

export default new Router({
    routes: [
        {
            path: '/',
            name: 'Mujer',
            component: Mujer
        },
        {
            path: '/agregar-embarazada',
            name: 'RegistrarEmbarazada',
            component: RegistrarEmbarazada,
        },


        {
            path: '/Login',
            name: 'Login',
            component: Login
        },
        {
            path: '/RegistrarPuerpera',
            name: 'RegistrarPuerpera',
            component: RegistrarPuerpera
        },
        {
            path: '/ActualizarEmbarazada',
            name: 'ActualizarEmbarazada',
            component: ActualizarEmbarazada,
        },
        {
            path: '/ActualizarPuerperia',
            name: 'ActualizarPuerperia',
            component: ActualizarPuerperia,
        },
        {
            path: '/ListaEmbarazada',
            name: 'ListaEmbarazada',
            component: ListaEmbarazada,
        },
        {
            path: '/ListaPuerpera',
            name: 'ListaPuerpera',
            component: ListaPuerpera,
        },
        {
            path: '/Informes',
            name: 'Informes',
            component: Informes,
        },
        {
            path: '/Tarjetas',
            name: 'Tarjetas',
            component: Tarjetas,
        },
        {
            path: '/ListaUsuario',
            name: 'ListaUsuario',
            component: ListaUsuario,
        },
    ]
});
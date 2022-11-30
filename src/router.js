import Vue from 'vue'
import Router from 'vue-router'
import Mujer from '@/components/Mujer'
import RegistrarEmbarazada from '@/components/RegistrarEmbarazada'
import Ajustes from '@/components/Ajustes'

import Informes from '@/components/Informes'

import Escritorio from '@/components/Escritorio'

import Login from '@/components/Login'
import RegistrarPuerpera from '@/components/RegistrarPuerpera'
import ActualizarEmbarazada from '@/components/ActualizarEmbarazada'
import ActualizarPuerperia from '@/components/ActualizarPuerperia'
import ListaPuerpera from '@/components/ListaPuerpera'
import ListaEmbarazada from '@/components/ListaEmbarazada'
import InformesRegional from '@/components/InformesRegional'
import InformesMunicipio from '@/components/InformesMunicipio'
import InformesLocalidad from '@/components/InformesLocalidad'

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
            path: '/ajustes',
            name: 'Ajustes',
            component: Ajustes,
        },
        {
            path: '/informes',
            name: 'Informes',
            component: Informes,
        },
        {
            path: '/escritorio',
            name: 'Escritorio',
            component: Escritorio,
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
            path: '/InformesRegional',
            name: 'InformesRegional',
            component: InformesRegional,
        },
        {
            path: '/InformesMunicipio',
            name: 'InformesMunicipio',
            component: InformesMunicipio,
        },
        {
            path: '/InformesLocalidad',
            name: 'InformesLocalidad',
            component: InformesLocalidad,
        },
    ]
});
<template>
    <section>
        <article class="panel centered">
            <p class="panel-heading">
                Bucar puerpera:
            </p>
            <b-field>
                <div class="control">
                    <b-button label="Eliminar usuario seleccionado" type="is-danger" icon-left="delete-forever"
                        :disabled="!selected" @click="selected = null" />
                </div>


                <b-button label="Modificar usuario seleccionado" type="is-success" icon-left="border-color"
                    :disabled="!selected" @click="selected = null" />
            </b-field>
        </article>
        <b-table id="listausuario" :data="usuarios" :columns="columns" :sticky-header="stickyHeaders" fixed-header height="450px"  :bordered="isbordered"></b-table>
    </section>
</template>
<script>
import DialogosService from "../services/DialogosService";
import LoginService from "../services/LoginService";

export default {
    data() {
        return {
            selected:null,
            isbordered: true,
            stickyHeaders: true,
            dateSearchable: false,
            usuarios: [],
            columns: [{
                field: "id",
                label: "No. Progreso",
                width: "40",
                centered: true,

            },
            {
                field: "clues",
                label: "Clues",
                centered: true,

            },
            {
                field: "nombreunidad",
                label: "Nombre Unidad",
                centered: true,
            },
            {
                field: "usuario",
                label: "Usuario",
                centered: true,
            },
            {
                field: "contraseña",
                label: "Contraseña",
                centered: true,
            },

            ],
        };
    },
    async mounted() {
        await this.ObtenerUsuarios();
    },
    methods: {

        async ObtenerUsuarios() {

            try {
                this.usuarios = await LoginService.obtenerUsuarios();
            } catch {
                DialogosService.mostrarNotificacionError("No se puede obtener la lista de usuarios");
            }
        }
    },


};
</script>
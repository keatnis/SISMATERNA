import { NotificationProgrammatic as Notification } from 'buefy'
import { DialogProgrammatic as Dialog } from 'buefy'


const DialogosService = {
    mostrarNotificacionExito(mensaje) {
        Notification.open({
            message: mensaje,
            type: 'is-info',

            position: "is-top-left"

        });
    },
    mostrarExitoso(mensaje) {
        Dialog.alert({
            title: 'Notificacion',
            message: mensaje,
            type: 'is-success',
            confirmText: 'Aceptar'
        })

    },
    mostrarError(message) {
        Dialog.alert({
            title: 'Error',
            message: message,
            type: 'is-danger',
            hasIcon: true,
            icon: 'times-circle',
            iconPack: 'fa',
            ariaRole: 'alertdialog',
            ariaModal: true
        })
    },
    mostrarNotificacionError(mensaje) {
        Notification.open({
            message: mensaje,
            type: 'is-danger',
            position: "is-top-left"
        });
    },
    mostrarDialogoConfirmacion(pregunta, callback, mensajeOk, mensajeCancelar) {
        mensajeOk = mensajeOk || "Ok";
        mensajeCancelar = mensajeCancelar || "Cancelar";
        Dialog.confirm({
            title: 'Confirmar',
            message: pregunta,
            cancelText: mensajeCancelar,
            confirmText: mensajeOk,
            type: 'is-danger',
            onConfirm: callback,
        });
    }
};
export default DialogosService;
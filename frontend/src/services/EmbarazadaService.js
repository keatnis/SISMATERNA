import HttpService from "./HttpService";
const VehiculosService = {
    async agregarVehiculo(vehiculo) {
        return await HttpService.post("/agregar_embarazada", vehiculo);
    },
    async obtenerEmbarazadas() {
        return await HttpService.get("/ListaEmbarazada");
    },
    async obtenerMunicipios(){
        return await HttpService.get("/agregar_embarazada");
       
    },
    async obteneLocalidadesById(idMunicipio) {
        return await HttpService.get(`/agregar_embarazada/${idMunicipio}`);
    },
    async agregarPuerpera(vehiculo) {
        return await HttpService.post("/agregar_embarazada", vehiculo);
    },

};
export default VehiculosService;
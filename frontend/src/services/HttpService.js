/*
const fetch = require("node-fetch");

fetch("https://reqres.in/api/users?page=2")
.then((respuesta) =>{
    return respuesta.json()
}).then((resp)=>{
    console.log(resp);
})
*/
 const RUTA_SERVIDOR = "http://localhost:8080"

const manejarRespuesta = async (respuesta) => {
    const respuestaDecodificada = await respuesta.json();
    if (respuestaDecodificada.error) {
        throw new Error(JSON.stringify(respuestaDecodificada.error));
    }
    console.log("respdecodificada ",respuestaDecodificada.data)
    return respuestaDecodificada.data;
    
};
const HttpService = {
    "post": async (ruta, datos) => {
        const respuestaRaw = await fetch(RUTA_SERVIDOR + ruta, {
            credentials: 'include',
            method: "POST",
            body: JSON.stringify(datos),
            mode: 'no-cors',

            
        });
        console.log("body post",JSON.stringify(datos));
        return await manejarRespuesta(respuestaRaw);
    },
    "put": async (ruta, datos) => {
        const respuestaRaw = await fetch(RUTA_SERVIDOR + ruta, {
            credentials: 'include',
            method: "PUT",
            body: JSON.stringify(datos)
        });
        return await manejarRespuesta(respuestaRaw);
    },
    "get": async (ruta) => {
        const respuestaRaw = await fetch(RUTA_SERVIDOR + ruta, {
            credentials: "include",
            mode: 'no-cors',
           
        });
       
        return await respuestaRaw.json();
     
    },
    "delete": async (ruta) => {
        const respuestaRaw = await fetch(RUTA_SERVIDOR + ruta, {
            credentials: 'include',
            method: 'DELETE',
        });
        return await manejarRespuesta(respuestaRaw);
    },
 };

 export default HttpService;


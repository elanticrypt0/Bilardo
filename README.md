# Bilardo

Bilardo es un web framework que utiliza Fiber y Gorm.
Ideal para crear APIs

## Diseño RFM (Routes Features Models)

### Routes

Las rutas se definen en la carpeta routes dentro del archivo setuproutes allí se agregan las funciones. Para poder acceder a las funcionalidades de Fiber o Gorm o la configuración de la app se debe pasar una variable

    bilardo *webcore.Bilardo

### Features

Las "features" son los controladores de las rutas y reciben 2 parámetros: 1) el contexto de Fiber y otra variable de *webcore.Bilardo para poder acceder a la db y configuración.

### Models

Los modelos deben recibir también la variable que apunte a *webcore.Bilardo


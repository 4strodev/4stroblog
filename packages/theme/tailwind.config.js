/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['../site/views/**/*.{html,js}', './src/**/*.{html,js}'],
    theme: {
        extend: {
            // Color palette https://coolors.co/2a6471-b99758-f7f2f7-c04c7b-7bd389
            colors: {
                'primary': {
                    DEFAULT: '#2a6471', 100: '#081416', 200: '#11282d', 300: '#193b43', 400: '#214f59', 500: '#2a6471', 600: '#3d91a4', 700: '#63b3c5', 800: '#97ccd8', 900: '#cbe6ec'
                },
                'secondary': {
                    DEFAULT: '#b99758', 100: '#271f10', 200: '#4e3e20', 300: '#745d31', 400: '#9b7c41', 500: '#b99758', 600: '#c8ad7b', 700: '#d5c19c', 800: '#e3d6bd', 900: '#f1eade'
                },
                'ghost_white': {
                    DEFAULT: '#f7f2f7', 100: '#3d253d', 200: '#794a79', 300: '#ad79ad', 400: '#d2b6d2', 500: '#f7f2f7', 600: '#f9f5f9', 700: '#faf7fa', 800: '#fcfafc', 900: '#fdfcfd'
                },
                'danger': {
                    DEFAULT: '#c04c7b', 100: '#280e18', 200: '#501c31', 300: '#782a49', 400: '#a03862', 500: '#c04c7b', 600: '#cd7196', 700: '#da95b0', 800: '#e6b8ca', 900: '#f3dce5'
                },
                'success': {
                    DEFAULT: '#7bd389', 100: '#113216', 200: '#21632c', 300: '#329542', 400: '#47c25c', 500: '#7bd389', 600: '#94dba0', 700: '#afe4b8', 800: '#c9edcf', 900: '#e4f6e7'
                }
            }
        },
    },
    plugins: [],
}


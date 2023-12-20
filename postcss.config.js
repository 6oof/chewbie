const cssnano = require('cssnano');
const tailwindcss = require('tailwindcss');

module.exports = {
  plugins: [
    tailwindcss,
    // other plugins...
    cssnano({
      preset: 'default',
    }),
  ],
};


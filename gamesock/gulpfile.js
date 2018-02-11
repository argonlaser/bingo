var gulp = require('gulp')
var nodemon = require('gulp-nodemon')

gulp.task('dev', () => {
  nodemon({
    script: 'server.js',
    ext: 'js',
    env: {
      'NODE_ENV': 'development'
    }
  })
})

gulp.task('prod', () => {
  nodemon({
    script: 'server.js',
    ext: 'js',
    env: {
      'NODE_ENV': 'production'
    }
  })
})

gulp.task('test', () => {
  nodemon({
    script: 'server.js',
    ext: 'js',
    env: {
      'NODE_ENV': 'test'
    }
  })
})

gulp.task('default', ['test'])

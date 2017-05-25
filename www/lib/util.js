module.exports = {
  view
}
function view(v){
  return `file://${__dirname}/../view/${v}.html`;
}
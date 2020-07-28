function loadFraction() {
  var request = $.ajax({
    url: "api/v1/fraction",
    method: "GET",
  });

  request.done(function( msg ) {
    katex.render(msg, fraction);
  });

  request.fail(function( jqXHR, textStatus ) {
    alert( "Request failed: " + textStatus );
  });
}
$( document ).ready(function() {
    console.log( "ready!" );
    loadFraction();
});

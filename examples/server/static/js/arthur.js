function loadFraction(fracType) {
  var request = $.ajax({
    url: "api/v1/fraction?type="+fracType,
    method: "GET",
  });

  request.done(function( msg ) {
    if (fracType == 'proper') {
      katex.render(msg, proper_fraction);
    } else if (fracType == 'improper') {
      katex.render(msg, improper_fraction);
    } else if (fracType == 'unit') {
      katex.render(msg, unit_fraction);
    }
  });

  request.fail(function( jqXHR, textStatus ) {
    alert( "Request failed: " + textStatus );
  });
}
function loadInteger(intType) {
  var request = $.ajax({
    url: "api/v1/integer?type="+intType,
    method: "GET",
  });

  request.done(function( msg ) {
    katex.render(msg, integer);
  });

  request.fail(function( jqXHR, textStatus ) {
    alert( "Request failed: " + textStatus );
  });
}
$( document ).ready(function() {
    console.log( "ready!" );
    loadInteger();
    loadFraction('proper');
    loadFraction('improper');
    loadFraction('unit');
});

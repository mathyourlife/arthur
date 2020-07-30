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
function loadInteger(size) {
  var request = $.ajax({
    url: "api/v1/integer?size="+size,
    method: "GET",
  });

  request.done(function( msg ) {
    if (size == 'small') {
      katex.render(msg, integer_small);
    } else if (size == 'medium') {
      katex.render(msg, integer_medium);
    } else if (size == 'large') {
      katex.render(msg, integer_large);
    }
  });

  request.fail(function( jqXHR, textStatus ) {
    alert( "Request failed: " + textStatus );
  });
}
function loadSum(sumType) {
  var request = $.ajax({
    url: "api/v1/sum",
    method: "GET",
  });

  request.done(function( msg ) {
    katex.render(msg, addition);
  });

  request.fail(function( jqXHR, textStatus ) {
    alert( "Request failed: " + textStatus );
  });
}
$( document ).ready(function() {
    console.log( "ready!" );
    loadInteger('small');
    loadInteger('medium');
    loadInteger('large');
    loadFraction('proper');
    loadFraction('improper');
    loadFraction('unit');
    loadSum('fraction');
});

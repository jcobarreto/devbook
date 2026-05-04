$('#login').on('submit', signin);

function signin(event) {
  event.preventDefault();

  $.ajax({
    url: '/login',
    method: 'POST',
    data: {
      email: $('#email').val(),
      password: $('#password').val()
    },
  }).done(function() {
    window.location = '/home';
  }).fail(function() {
    Swal.fire("Ops...", "Invalid email or password!", "error");
  });
}

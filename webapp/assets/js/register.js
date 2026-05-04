$('#signup-form').on('submit', createUser);

function createUser(event) {
  event.preventDefault();

  if ($('#password').val() !== $('#confirm-password').val()) {
    Swal.fire({
      title: 'Error',
      text: 'Passwords do not match!',
      icon: 'error',
      confirmButtonText: 'OK'
    });
    return;
  }

  $.ajax({
    url: '/users',
    method: 'POST',
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      password: $('#password').val()
    },
    dataType: 'text',
  }).done(function() {
    Swal.fire({
      title: 'Success',
      text: 'User created successfully!',
      icon: 'success',
      confirmButtonText: 'OK'
    }).then(() => {
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
        Swal.fire("Ops...", "Error logging in after registration!", "error");
      });
    });
  }).fail(function() {
    Swal.fire({
      title: 'Error',
      text: 'Error creating user!',
      icon: 'error',
      confirmButtonText: 'OK'
    });
  });
}

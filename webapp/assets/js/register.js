$('#signup-form').on('submit', createUser);

function createUser(event) {
  event.preventDefault();

  if ($('#password').val() !== $('#confirm-password').val()) {
    alert('Passwords do not match. Please try again.');
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
  }).done(function() {
    alert('User created successfully!');
  }).fail(function() {
    alert('Error creating user!');
  });
}

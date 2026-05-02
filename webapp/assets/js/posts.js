$('#new-post').on('submit', createPost);

$(document).on('click', '.like-post', likePost);
$(document).on('click', '.unlike-post', unlikePost);

$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost);

function createPost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    },
    dataType: "text",
  }).done(function() {
    window.location = '/home';
  }).fail(function() {
    Swal.fire("Ops!", "Error creating post!", "error");
  });
}

function likePost(event) {
  event.preventDefault();

  const clickedElement = $(event.currentTarget);
  const postID = clickedElement.closest('.p-5').data('post-id');

  clickedElement.prop('disabled', true);
  $.ajax({
    url: `/posts/${postID}/like`,
    method: "POST",
  }).done(function() {
    const likeCountElement = clickedElement.next('span');
    const likeCount = parseInt(likeCountElement.text());

    likeCountElement.text(likeCount + 1);

    clickedElement.addClass('unlike-post')
    clickedElement.addClass('text-danger')
    clickedElement.removeClass('like-post');

  }).fail(function() {
    Swal.fire("Ops!", "Error liking post!", "error");
  }).always(function() {
    clickedElement.prop('disabled', false);
  });
}

function unlikePost(event) {
  event.preventDefault();

  const clickedElement = $(event.currentTarget);
  const postID = clickedElement.closest('.p-5').data('post-id');

  clickedElement.prop('disabled', true);
  $.ajax({
    url: `/posts/${postID}/unlike`,
    method: "POST",
  }).done(function() {
    const likeCountElement = clickedElement.next('span');
    const likeCount = parseInt(likeCountElement.text());

    likeCountElement.text(likeCount - 1);

    clickedElement.removeClass('unlike-post');
    clickedElement.removeClass('text-danger')
    clickedElement.addClass('like-post')

  }).fail(function() {
    Swal.fire("Ops!", "Error unliking post!", "error");
  }).always(function() {
    clickedElement.prop('disabled', false);
  });
}

function updatePost() {
  $(this).prop('disabled', true);

  const postID = $(this).data('post-id');

  $.ajax({
    url: `/posts/${postID}`,
    method: "PUT",
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    }
  }).done(function() {
    Swal.fire({
      title: 'Post Updated',
      text: 'Your post has been updated successfully.',
      icon: 'success',
      confirmButtonText: 'OK'
    }).then(() => {
      window.location = '/home';
    });
  }).fail(function() {
    Swal.fire("Ops!", "Error updating post!", "error");
  }).always(function() {
    $('#update-post').prop('disabled', false);
  });
}

function deletePost(event) {
  event.preventDefault();

  Swal.fire({
    title: 'Are you sure?',
    text: "This action cannot be undone!",
    icon: 'warning',
    showCancelButton: true,
    confirmButtonColor: '#d33',
    cancelButtonColor: '#3085d6',
    confirmButtonText: 'Yes, delete it!'
  }).then((result) => {
    if (result.isConfirmed) {
      const clickedElement = $(event.target);
      const post = clickedElement.closest('div');
      const postID = post.data('post-id');

      $.ajax({
        url: `/posts/${postID}`,
        method: "DELETE",
      }).done(function() {
        post.fadeOut("slow", function() {
          $(this).remove();
        });
      }).fail(function() {
        Swal.fire("Ops!", "Error deleting post!", "error");
      });
    }
  });
}

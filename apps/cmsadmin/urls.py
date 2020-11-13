from django.urls import path

from . import views, user_view, meetplan_view

app_name = 'cmsadmin'

urlpatterns = [
    path('', views.noindex),
    path('index/', views.IndexView.as_view(), name="index"),
]

urlpatterns += [
    path('user_teacher_all/', user_view.TeacherListView.as_view(), name='user_teacher_all'),
    path('user_student_all/', user_view.StudentListView.as_view(), name='user_student_all'),
    path('user_create/', user_view.UserCreateView.as_view(), name='user_create'),
    path('user_create_many/', user_view.CreateManyUserView.as_view(), name='user_create_many'),
    path('user_detail/<int:pk>/', user_view.UserDetailView.as_view(), name='user-detail'),
    path('user_update/<int:pk>/', user_view.UpdateUserView.as_view(), name='user_update'),
    path('user_deletelist/', user_view.DeletedUserListView.as_view(), name='user_delete_list'),
    path('user_recovery/<int:pk>/', user_view.RecoveryUserView.as_view(), name='user-recovery'),
    path('user_confirm_delete/<int:pk>/', user_view.UserDeleteView.as_view(), name='user-confirm-delete'),

    path('user_base_profile_update/<int:pk>/', user_view.BaseProfileUpdateView.as_view(), name='base-profile-update'),
    path('user_student_profile_update/<int:pk>/', user_view.StudentProfileUpdateView.as_view(),
         name='student-profile-update'),
    path('user_teacher_profile_update/<int:pk>/', user_view.TeacherProfileUpdateView.as_view(),
         name='teacher-profile-update'),

    path('user_management/', user_view.OtherManagementView.as_view(), name='user-other-management'),
    path('user_major_all/', user_view.MajorListView.as_view(), name='user-major-all'),
    path('user_major_create/', user_view.MajorCreateView.as_view(), name='user-major-create'),
    path('user_major_update/<int:pk>/', user_view.MajorUpdateView.as_view(), name='user-major-update'),
    path('user_major_delete/<int:pk>/', user_view.MajorDeleteView.as_view(), name='user-major-delete'),
    path('user_department_all/', user_view.DepartmentListView.as_view(), name='user-department-all'),
    path('user_department_create/', user_view.DepartmentCreateView.as_view(), name='user-department-create'),
    path('user_department_update/<int:pk>/', user_view.DepartmentUpdateView.as_view(), name='user-department-update'),
    path('user_department_delete/<int:pk>/', user_view.DepartmentDeleteView.as_view(), name='user-department-delete'),
    path('user_grade_all/', user_view.GradeListView.as_view(), name='user-grade-all'),
    path('user_grade_create/', user_view.GradeCreateView.as_view(), name='user-grade-create'),
    path('user_grade_update/<int:pk>/', user_view.GradeUpdateView.as_view(), name='user-grade-update'),
    path('user_grade_delete/<int:pk>/', user_view.GradeDeleteView.as_view(), name='user-grade-delete'),

]

urlpatterns += [
    path('meetplan_all/', meetplan_view.MeetPlanListView.as_view(), name='meetplan_all'),
    path('meetplan_create/', meetplan_view.MeetPlanCreateView.as_view(), name='meetplan_create'),
    path('meetplan_create/<int:pk>/', meetplan_view.MeetPlanCreateFromTeacherView.as_view(),
         name='meetplan-create-from-teacher'),
    path('meetplan_update/<int:pk>/', meetplan_view.MeetPlanUpdateView.as_view(), name='meetplan_update'),
    path('meetplan_detail/<int:pk>/', meetplan_view.MeetPlanDetailView.as_view(), name='meetplan-detail'),
    path('meetplan_confirm_delete/<int:pk>/', meetplan_view.MeetPlanDeleteView.as_view(),
         name='meetplan-confirm-delete'),

    path('meetplanorder_all/', meetplan_view.MeetPlanOrderListView.as_view(), name='meetplanorder_all'),
    path('meetplanorder_create/', meetplan_view.MeetPlanOrderCreateView.as_view(), name='meetplanorder_create'),
    path('meetplanorder_create/<int:pk>/', meetplan_view.MeetPlanOrderCreateFromStudentView.as_view(),
         name='meetplanorder-create-from-student'),
    path('meetplanorder_update/<int:pk>/', meetplan_view.MeetPlanOrderUpdateView.as_view(),
         name='meetplanorder_update'),
    path('meetplanorder_confirm_delete/<int:pk>/', meetplan_view.MeetPlanOrderDeleteView.as_view(),
         name='meetplanorder-confirm-delete'),

    path('meetplanorder_create_many/', meetplan_view.MeetPlanUndergraduateResearch.as_view(),
         name='meetplanorder-undergraduate-research'),

    path('meetplan_feedback_all/', meetplan_view.FeedBackListView.as_view(), name='feedback_all'),
    path('meetplan_feedback_update/<int:pk>/', meetplan_view.FeedBackUpdateView.as_view(), name='feedback_update'),

    path('meetplan_report/', meetplan_view.MeetPlanReportListView.as_view(), name='meetplan-report-all'),
    path('meetplan_report_teacher_create/', meetplan_view.MeetPlanReportTeacherCreateView.as_view(),
         name='meetplan-report-teacher-create'),
    path('meetplan_report_student_create/', meetplan_view.MeetPlanReportStudentCreateView.as_view(),
         name='meetplan-report-student-create'),
    path('meetplan_term_date_setting/', meetplan_view.TermDateUpdateView.as_view(), name='meetplan-termdate-update')
]

urlpatterns += [
    path('site_updaterecord_all/', views.UpdateRecordListView.as_view(), name='site-updaterecord-all'),
    path('site_updaterecord_create/', views.UpdateRecordCreateView.as_view(), name='site-updaterecord-create'),
    path('site_updaterecord_update/<int:pk>/', views.UpdateRecordUpdateView.as_view(), name='site-updaterecord-update'),
    path('site_updaterecord_delete/<int:pk>/', views.UpdateRecordDeleteView.as_view(), name='site-updaterecord-delete'),
]

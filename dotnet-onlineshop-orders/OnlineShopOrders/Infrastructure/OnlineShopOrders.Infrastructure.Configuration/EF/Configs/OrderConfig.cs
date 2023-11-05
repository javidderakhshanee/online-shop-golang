using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

using OnlineShopOrders.Core.Domain;

namespace OnlineShopOrders.Infrastructure.Configuration.EF.Configs;

public class OrderConfig : IEntityTypeConfiguration<Order>
{
    public void Configure(EntityTypeBuilder<Order> builder)
    {
        builder.ToTable("orders");

        builder.HasKey(x => x.Id);

        builder.Property(x => x.CustomerId)
            .IsRequired();

        builder.Property(x => x.AddressId)
            .IsRequired();

        builder.Property(x => x.CreatedDate)
            .IsRequired();

        builder.HasMany(t => t.OrderLines)
        .WithOne(t => t.Order)
        .HasForeignKey(d => d.OrderId)
        .HasPrincipalKey(e => e.Id)
        .OnDelete(DeleteBehavior.Cascade);
    }
}
